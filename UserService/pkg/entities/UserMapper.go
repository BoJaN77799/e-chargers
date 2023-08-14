package entities

func (user *User) ToDTO() UserReservationDTO {
	return UserReservationDTO{
		Vehicles: vehiclesToDto(user.Vehicles),
	}
}
func (user *User) ToUserProfileDTO() UserProfileDTO {
	return UserProfileDTO{
		Id:        user.Id,
		Email:     user.Email,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Strikes:   user.Strikes,
		Vehicles:  vehiclesToDto(user.Vehicles),
	}
}

func (user *User) ToReportDTO() UserReportDTO {
	return UserReportDTO{
		Id:          user.Id,
		Email:       user.Email,
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Role:        user.Role.String(),
		Strikes:     user.Strikes,
		Banned:      user.Banned,
		BannedAt:    user.BannedAt,
		BannedUntil: user.BannedUntil,
	}
}
func (vehicle *Vehicle) ToDTO() VehicleDto {
	return VehicleDto{
		Id:          vehicle.Id,
		Name:        vehicle.Name,
		VehicleType: vehicle.VehicleType.String(),
	}
}

func vehiclesToDto(vehicles []Vehicle) []VehicleDto {
	var vehiclesDTO []VehicleDto
	for _, vehicle := range vehicles {
		vehiclesDTO = append(vehiclesDTO, vehicle.ToDTO())
	}
	return vehiclesDTO
}

func (vehicle *VehicleDto) FromDto() Vehicle {
	return Vehicle{
		Name:        vehicle.Name,
		VehicleType: StrToVehicleType(vehicle.VehicleType),
	}
}
