import { HttpResponse } from '@angular/common/http';
import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ChargerDTO } from 'src/modules/shared/model/chargerDTO';
import { SnackBarService } from 'src/modules/shared/service/snack-bar.service';
import { SearchDTO } from 'src/modules/user/model/searchDTO';
import { ChargerService } from 'src/modules/user/service/chargerService';

@Component({
  selector: 'app-search-chargers',
  templateUrl: './search-chargers.component.html',
  styleUrls: ['./search-chargers.component.scss']
})
export class SearchChargersComponent {

  @Output()
  searchedChargersEvent = new EventEmitter<ChargerDTO[]>();

  public searchFormGroup: FormGroup;

  constructor(private formBuilder: FormBuilder, private snackBarService: SnackBarService, private chargerService: ChargerService) {
    this.searchFormGroup = this.formBuilder.group({
      name: ['Promenada'],
      workTimeFrom: [0, [Validators.min(0), Validators.max(24)]],
      workTimeTo: [24, [Validators.min(0), Validators.max(24)]],
      capacity: [10],
      pricePerHourFrom: [1, [Validators.min(1), Validators.max(20)]],
      pricePerHourTo: [20, [Validators.min(1), Validators.max(20)]],
      type: ['Type 2', Validators.required],
      chargingSpeedFrom: [5, [Validators.min(5), Validators.max(35)]],
      chargingSpeedTo: [35, [Validators.min(5), Validators.max(35)]],
    });
  }

  get searchForm() {
    return this.searchFormGroup.value
  }

  submit() {
    try {
      this.validate();
    } catch (e: any) {
      this.snackBarService.openSnackBar(e.message)
      return
    }

    this.chargerService.search({ ...this.searchForm }).subscribe(
      (res) => this.onSuccess(res),
      (err) => this.snackBarService.openSnackBar(err.error)
    )
  }

  onSuccess(res: HttpResponse<ChargerDTO[]>) {
    if (!res.body) {
      this.snackBarService.openSnackBar("Search with given inputs didnt find any charger.")
      return
    }
    this.searchedChargersEvent.emit(res.body)
  }

  validate() {
    const isValidWorkTime = this.searchForm.workTimeFrom < this.searchForm.workTimeTo;
    if (!isValidWorkTime) throw new Error('Invalid work time entered.')

    const isValidPrice = this.searchForm.pricePerHourFrom < this.searchForm.pricePerHourTo;
    if (!isValidPrice) throw new Error('Invalid price entered.')

    const isValidChargingSpeed = this.searchForm.chargingSpeedFrom < this.searchForm.chargingSpeedTo;
    if (!isValidChargingSpeed) throw new Error('Invalid charging speed entered')
  }
}

