package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	app "github.com/ithinkiborkedit/GUSH-Client.git/application/client"
	dclient "github.com/ithinkiborkedit/GUSH-Client.git/domain/client"
	"github.com/ithinkiborkedit/GUSH-Client.git/infrastructure/netclient"
	"github.com/ithinkiborkedit/GUSH-Client.git/infrastructure/storage"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s server-address player-name\n", os.Args[0])
	}

	serverAddr := os.Args[1]
	playerName := strings.Join(os.Args[2:], " ")

	playerRepo := storage.NewInMemoryPlayerRepo()
	worldRepo := storage.NewInMemoryWorldRepo()

	renderer := &dclient.RendererService{}

	tcpClient := netclient.NewTCPNetClient()

	useCase := app.ClientUseCase{
		PlayerRepo: playerRepo,
		WorldRepo:  worldRepo,
		Renderer:   renderer,
		NetClient:  tcpClient,
	}

	err := useCase.ConnectToServer(serverAddr, "local-player", playerName)
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	fmt.Printf("connected to %s as %s\n", serverAddr, playerName)

	useCase.ListenAsync()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading input exiting")
			break
		}
		cmdLine := strings.TrimSpace(line)
		if cmdLine == "/quit" {
			_ = tcpClient.Close()
			fmt.Println("Disconnected")
			break
		}

		if cmdLine == "" {
			continue

		}

		parts := strings.SplitN(cmdLine, " ", 2)
		cmdType := parts[0]
		payload := ""
		if len(parts) > 1 {
			payload = parts[1]
		}

		err = useCase.SendCommand(cmdType, payload)
		if err != nil {
			fmt.Println("error sending command")
		}
	}
}
