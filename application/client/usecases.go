package client

import (
	"sync"

	"github.com/ithinkiborkedit/GUSH-Client.git"
	dclient "github.com/ithinkiborkedit/GUSH-Client.git/domain/client"
)

type ClientUseCase struct {
	PlayerRepo dclient.PlayerRepository
	WorldRepo  dclient.WorldRepository
	Renderer   *dclient.RendererService
	NetClient  NetClient

	mu sync.Mutex
}

type NetClient interface {
	Connect(address string) error
	SendCommand(cmd *GUSH.Command) error
	ReadLoop(callback func(*GUSH.ServerMessage, error))
	Close() error
}

func (uc *ClientUseCase) ConnectToServer(address, playerID, playerName string) error {
	err := uc.NetClient.Connect(address)
	if err != nil {
		return err
	}

	player := &dclient.LocalPlayer{
		ID:   playerID,
		Name: playerName,
	}

	err = uc.PlayerRepo.SaveLocalPlayer(player)
	if err != nil {
		return err
	}

	world, err := uc.WorldRepo.GetWorld()
	if err != nil {
		return err
	}

	world.Player = player
	_ = uc.WorldRepo.SaveWorld(world)

	return nil
}

func (uc *ClientUseCase) ListenAsync() {
	go uc.NetClient.ReadLoop(func(msg *GUSH.ServerMessage, err error) {
		if err != nil {
			return
		}
		if msg != nil {
			uc.handleServerMessage(msg)
		}
	})
}

func (uc *ClientUseCase) handleServerMessage(msg *GUSH.ServerMessage) {
	uc.mu.Lock()
	defer uc.mu.Unlock()

	world, _ := uc.WorldRepo.GetWorld()
	switch payload := msg.Payload.(type) {
	case *GUSH.ServerMessage_Chat:
		chatStr := uc.Renderer.RenderChat(payload.Chat.PlayerName, payload.Chat.Text)
		print(chatStr)

	case *GUSH.ServerMessage_RoomUpdate:
		r := &dclient.LocalRoom{
			ID:          payload.RoomUpdate.RoomId,
			Name:        payload.RoomUpdate.RoomName,
			Description: payload.RoomUpdate.Description,
		}
		world.Rooms[r.ID] = r
		if world.Player != nil {
			world.Player.Room = r.ID
		}
		_ = uc.WorldRepo.SaveWorld(world)
		roomStr := uc.Renderer.RenderRoomUpdate(r.ID, r.Name, r.Description)
		print(roomStr)

	case *GUSH.ServerMessage_SystemMsg:
		sysStr := uc.Renderer.RenderSystemMsg(payload.SystemMsg.Text)
		print(sysStr)
	}
}

func (uc *ClientUseCase) SendCommand(cmdType, payload string) error {
	cmd := &GUSH.Command{
		Type:    cmdType,
		Payload: payload,
	}
	return uc.NetClient.SendCommand(cmd)
}
