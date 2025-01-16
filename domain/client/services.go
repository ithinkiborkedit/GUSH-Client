package client

import "fmt"

type RendererService struct{}

func (r *RendererService) RenderChat(playerName, text string) string {
	return fmt.Sprintf("%s says %s\n", playerName, text)
}

func (r *RendererService) RenderSystemMsg(text string) string {
	return fmt.Sprintf("[SYSTEM] %s\n", text)
}

func (r *RendererService) RenderRoomUpdate(roomID, name, desc string) string {
	return fmt.Sprintf("\n[ROOM: %s]\nName:%s\nDescription: %s\n\n", roomID, name, desc)
}
