import { Component, OnInit, ViewChild } from '@angular/core';
import { ChargerInfoComponent } from 'src/modules/shared/components/charger-info/charger-info.component';
import { ChargerDTO } from 'src/modules/shared/model/chargerDTO';
import { MapPageComponent } from 'src/modules/shared/pages/map-page/map-page.component';
import { SnackBarService } from 'src/modules/shared/service/snack-bar.service';
import { ChargerService } from 'src/modules/user/service/chargerService';

@Component({
  selector: 'app-admin-home',
  templateUrl: './admin-home.component.html',
  styleUrls: ['./admin-home.component.scss']
})
export class AdminHomeComponent implements OnInit {

  chargers: ChargerDTO[]
  chargersToMap: ChargerDTO[]

  selectedCharger: ChargerDTO | undefined

  @ViewChild(MapPageComponent)
  child!: MapPageComponent;

  @ViewChild(ChargerInfoComponent) chargerInfoComponent!: ChargerInfoComponent;

  // search
  searchOpened: boolean = false

  constructor(private chargerService: ChargerService, private snackBarService: SnackBarService) {
    this.chargers = []
    this.chargersToMap = []
  }

  ngOnInit(): void {
    this.chargerService.getAllChargers().subscribe(
      (response) => {
        this.chargers = response.body as ChargerDTO[]
        console.log(this.chargers)
        this.selectedCharger = this.chargers[0]
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
    this.searchOpened = false;
    this.chargerInfoComponent.recensions = [];
    this.chargerInfoComponent.reviewsVisible = false;
  }

  toggleSearch() {
    this.searchOpened = !this.searchOpened;
  }

  changeChargers(chargers: ChargerDTO[]) {
    console.log(chargers)
    this.chargersToMap = chargers;
    this.child.initMap(chargers);
  }

}
