import { Component, OnInit } from '@angular/core';
import { Coordinates } from 'src/modules/shared/model/coordinates';
import { ChargerDTO } from 'src/modules/shared/model/chargerDTO';
import { ChargerService } from '../../service/chargerService';

@Component({
  selector: 'app-user-home',
  templateUrl: './user-home.component.html',
  styleUrls: ['./user-home.component.scss']
})
export class UserHomeComponent implements OnInit {

  chargers: ChargerDTO[]
  chargersToMap: ChargerDTO[]

  selectedCharger: ChargerDTO | undefined

  // search
  searchOpened: boolean = true

  constructor(private chargerService: ChargerService) {
    this.chargers = []
    this.chargersToMap = []
  }

  ngOnInit(): void {
    this.loadChargers();
  }

  loadChargers() {
    this.chargerService.getAllChargers().subscribe(
      (response) => {
        console.log(response.body)
        this.chargers = response.body as ChargerDTO[]
        this.chargers.forEach(chargerForMap => {
          this.chargersToMap.push(chargerForMap)
        });
      },
      (err) => {
        console.log(err.error)
      }
    )
  }

  setSelectedCharger(charger: ChargerDTO) {
    this.selectedCharger = charger;
  }

  toggleSearch() {
    this.searchOpened = !this.searchOpened;
  }

}
