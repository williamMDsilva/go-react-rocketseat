export interface CreateMessageReactionRequest {
    roomId: string
    messageId: string
}

export interface CreateMessageRequest {
    roomId: string
    message: string
}

export interface CreateRoomRequest {
    theme: string
}

export interface GetRoomMessagesRequest {
    roomId: string
}

export interface RemoveMessageReactionRequest {
    roomId: string
    messageId: string
}