import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Charger } from 'src/modules/shared/model/chargerDTO';
import { SnackBarService } from 'src/modules/shared/service/snack-bar.service';
import { UtilService } from 'src/modules/shared/service/utils-service';
import { ChargerService } from 'src/modules/user/service/chargerService'; import { VehicleService } from 'src/modules/user/service/vehicleService';

@Component({
  selector: 'app-create-charger',
  templateUrl: './create-charger.component.html',
  styleUrls: ['./create-charger.component.scss']
})
export class CreateChargerComponent implements OnInit {

  reservationForm: FormGroup;

  coordinatesGetted = false;

  constructor(private chargerService: ChargerService,
    private snackBarService: SnackBarService,
    private fb: FormBuilder, private router: Router,
    private utilService: UtilService
  ) {
    this.reservationForm = this.fb.group({
      name: ["Dom B punjac", [Validators.required]],
      street: ["Bulevar Despota Stefana 7", [Validators.required]],
      city: ["Novi Sad", [Validators.required]],
      country: ["Srbija", [Validators.required]],
      postal_code: [21000, [Validators.required]],
      work_time_from: [0, [Validators.required, Validators.min(0), Validators.max(24)]],
      work_time_to: [24, [Validators.required, Validators.min(0), Validators.max(24)]],
      capacity: [7, [Validators.required, Validators.min(1), Validators.max(10)]],
      price_per_hour: [7, [Validators.required, Validators.min(1), Validators.max(20)]],
      charging_speed: [25, [Validators.required, Validators.min(5), Validators.max(35)]],
    });
  }


  ngOnInit(): void {

  }

  submit() {
    let charger: Charger = {
      name: this.reservationForm.get("name")?.value,
      address: {
        street: this.reservationForm.get("street")?.value,
        city: this.reservationForm.get("city")?.value,
        country: this.reservationForm.get("country")?.value,
        latitude: 0,
        longitude: 0,
        postal_code: this.reservationForm.get("postal_code")?.value
      },
      work_time_from: this.reservationForm.get("work_time_from")?.value,
      work_time_to: this.reservationForm.get("work_time_to")?.value,
      capacity: this.reservationForm.get("capacity")?.value,
      description: "WiFi, Food, Shopping, Free Parking",
      rating: 0,
      plugs: "Type 1, Type 2",
      price_per_hour: this.reservationForm.get("price_per_hour")?.value,
      charging_speed: this.reservationForm.get("charging_speed")?.value,
    }

    // getting coordinates

    this.utilService.findCoordinated(charger.address).subscribe(
      (response) => {
        console.log(response)
        if (response.body && response.body.lenght != 0) {
          let address = response.body[0]
          console.log(address)
          console.log(address.lon)
          console.log(address.lat)
          charger.address.longitude = Number(address.lon)
          charger.address.latitude = Number(address.lat)
          console.log(charger)

          // let's save this in database
          this.chargerService.create(charger).subscribe(
            (response) => {
              console.log(response)
              if (response.body)
                this.snackBarService.openSnackBar(response.body)
              this.router.navigate(["/myapp/admin/homepage"]);
            },
            (err) => {
              this.snackBarService.openSnackBarFast(err.error)
            }
          )
        }
      },
      (err) => {
        console.log(err.error)
        this.snackBarService.openSnackBar("Error getting coordinates")
      }
    )
  }

}
