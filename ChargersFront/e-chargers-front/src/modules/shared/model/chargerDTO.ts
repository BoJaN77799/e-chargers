export interface AdressDTO {
    street: string
    city: string
    country: string
    postal_code: number
    latitude: number
    longitude: number
}


export interface ChargerDTO {
    id: number
    name: string
    address: AdressDTO
    work_time: string,
    capacity: number,
    description: string,
    plugs: string[],
    price_per_hour: string,
    charging_speed: string
}

export interface Address {
    street: string
    city: string
    country: string
    postal_code: number
    longitude: number
    latitude: number
}

export interface Charger {
    name: string,
    address: Address
    work_time_from: number
    work_time_to: number
    capacity: number
    description: string
    rating: number
    plugs: string
    price_per_hour: number
    charging_speed: number
}