import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { SnackBarService } from 'src/modules/shared/service/snack-bar.service';
import { SearchDTO } from '../../model/searchDTO';

@Component({
  selector: 'app-search-chargers',
  templateUrl: './search-chargers.component.html',
  styleUrls: ['./search-chargers.component.scss']
})
export class SearchChargersComponent implements OnInit {

  public searchFormGroup: FormGroup;

  constructor(private fb: FormBuilder, private snackBarService: SnackBarService) {
    this.searchFormGroup = this.fb.group({
      searchField: ['', Validators.required],
      workTimeFrom: ['00'],
      workTimeTo: [24],
      capacity: [10],
      pricePerHourFrom: [50],
      pricePerHourTo: [700],
      type: ['', Validators.required],
      chargingSpeedFrom: [5],
      chargingSpeedTo: [35],
    });
  }

  ngOnInit(): void {
  }

  submit() {

    if (!this.validWorkTime())
      return
    if (!this.validPrice())
      return
    if (!this.validChargingSpeed())
      return

    let searchDTO: SearchDTO = {
      "searchField": this.searchFormGroup.get('searchField')?.value,
      "workTimeFrom": Number(this.searchFormGroup.get('workTimeFrom')?.value),
      "workTimeTo": this.searchFormGroup.get('workTimeTo')?.value,
      "capacity": this.searchFormGroup.get('capacity')?.value,
      "pricePerHourFrom": this.searchFormGroup.get('pricePerHourFrom')?.value,
      "pricePerHourTo": this.searchFormGroup.get('pricePerHourTo')?.value,
      "type": this.searchFormGroup.get('type')?.value,
      "chargingSpeedFrom": this.searchFormGroup.get('chargingSpeedFrom')?.value,
      "chargingSpeedTo": this.searchFormGroup.get('chargingSpeedTo')?.value,
    }

    console.log(searchDTO);

    this.snackBarService.openSnackBar("Sve dobro")

  }

  validWorkTime(): boolean {
    if (Number(this.searchFormGroup.get('workTimeFrom')?.value) > this.searchFormGroup.get('workTimeTo')?.value) {
      this.snackBarService.openSnackBar("Work time - From must be lower than To")
      return false;
    }
    return true;
  }

  validPrice(): boolean {
    if (this.searchFormGroup.get('pricePerHourFrom')?.value > this.searchFormGroup.get('pricePerHourTo')?.value) {
      this.snackBarService.openSnackBar("Price - From must be lower than To")
      return false;
    }
    return true;
  }

  validChargingSpeed(): boolean {
    if (this.searchFormGroup.get('chargingSpeedFrom')?.value > this.searchFormGroup.get('chargingSpeedTo')?.value) {
      this.snackBarService.openSnackBar("Charging speed - From must be lower than To")
      return false;
    }
    return true;
  }

}

