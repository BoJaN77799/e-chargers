export interface RecensionDTO {
    id: number
    user_id: string
    charger_id: number
    date: number
    content: string
    rate: number,
    toxic: number
}

export type RecensionWithUserDTO = RecensionDTO & {
    banned: boolean,
    email: string,
    firstname: string,
    lastname: string
} 