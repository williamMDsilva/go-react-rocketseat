export interface GetRoomMessagesResponse {
    messages: {
        id: string;
        text: string;
        amountOfReactions: number;
        answered: boolean;
    }[]
}