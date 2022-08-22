import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatTableDataSource } from '@angular/material/table';
import { SnackBarService } from 'src/modules/shared/service/snack-bar.service';
import { UtilService } from 'src/modules/shared/service/utils-service';
import { VehicleDTO } from 'src/modules/user/model/vehicleDTO';
import { VehicleService } from '../../service/vehicleService';

@Component({
  selector: 'app-vehicles',
  templateUrl: './vehicles.component.html',
  styleUrls: ['./vehicles.component.scss']
})
export class VehiclesComponent implements OnInit {

  displayedColumns: string[] = ['name', 'type', 'button'];
  dataSource: MatTableDataSource<VehicleDTO>;
  vehicles: VehicleDTO[] = []

  addVehicleForm: FormGroup


  constructor(private fb: FormBuilder, private vehiclesService: VehicleService, private utilService: UtilService, private snackBarService: SnackBarService) {
    this.dataSource = new MatTableDataSource(this.vehicles)
    this.addVehicleForm = this.fb.group({
      name: ['', Validators.required],
      vehicle_type: ['', Validators.required],
    });
  }

  ngOnInit(): void {

    this.vehiclesService.getAllVehiclesOfUser(this.utilService.getLoggedUsername()).subscribe(
      (response) => {
        if (response) {
          this.vehicles = response.body as VehicleDTO[];
          console.log(this.vehicles)
          this.dataSource = new MatTableDataSource(this.vehicles);
        }
      },
      (err) => {
        this.snackBarService.openSnackBarFast(err.error)
      }
    )
  }

  submit() {

    let vehicle: VehicleDTO = {
      "name": this.addVehicleForm.get("name")?.value,
      "vehicle_type": this.addVehicleForm.get("vehicle_type")?.value,
      "username": this.utilService.getLoggedUsername(),
      "id": 0,
    }

    this.vehiclesService.save(vehicle).subscribe(
      (response) => {
        if (response) {
          this.vehicles.push(vehicle);
          this.dataSource = new MatTableDataSource(this.vehicles);
        }
      },
      (err) => {
        this.snackBarService.openSnackBarFast(err.error)
      }
    )

  }

  remove(name: string) {
    this.vehiclesService.remove(name).subscribe(
      (response) => {
        if (response) {
          this.vehicles.splice(this.vehicles.findIndex((element) => element.name === name));
          this.dataSource = new MatTableDataSource(this.vehicles);
        }
      },
      (err) => {
        this.snackBarService.openSnackBarFast(err.error)
      }
    )
  }

}