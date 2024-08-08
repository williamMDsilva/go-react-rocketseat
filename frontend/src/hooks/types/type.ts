export type WebhookMessage =
    | { kind: "message_created"; value: { id: string, message: string } }
    | { kind: "message_answered"; value: { id: string } }
    | { kind: "message_reaction_increased"; value: { id: string; count: number } }
    | { kind: "message_reaction_decreased"; value: { id: string; count: number } };
