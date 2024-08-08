import { GetRoomMessagesResponse } from "../types/interfaces"
import { GetRoomMessagesRequest } from "./types/interfaces"


export async function getRoomMessages({ roomId }: GetRoomMessagesRequest): Promise<GetRoomMessagesResponse> {
  const response = await fetch(`${import.meta.env.VITE_APP_API_URL}/rooms/${roomId}/messages`)

  const data: Array<{
    ID: string
    RoomId: string
    Message: string
    ReactionCount: number
    Answered: boolean
  }> = await response.json()

  return {
    messages: data.map(item => {
      return {
        id: item.ID,
        text: item.Message,
        amountOfReactions: item.ReactionCount,
        answered: item.Answered,
      }
    })
  }
}