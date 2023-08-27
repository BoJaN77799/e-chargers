import { Component, OnInit, ViewChild } from '@angular/core';
import { Coordinates } from 'src/modules/shared/model/coordinates';
import { ChargerDTO } from 'src/modules/shared/model/chargerDTO';
import { ChargerService } from '../../service/chargerService';
import { ChargerInfoComponent } from 'src/modules/shared/components/charger-info/charger-info.component';
import { MapPageComponent } from 'src/modules/shared/pages/map-page/map-page.component';

@Component({
  selector: 'app-user-home',
  templateUrl: './user-home.component.html',
  styleUrls: ['./user-home.component.scss']
})
export class UserHomeComponent implements OnInit {

  chargers: ChargerDTO[]
  chargersToMap: ChargerDTO[]

  selectedCharger: ChargerDTO

  // search
  searchOpened: boolean = false

  @ViewChild(MapPageComponent) mapComponent!: MapPageComponent;


  @ViewChild(ChargerInfoComponent) chargerInfoComponent!: ChargerInfoComponent;

  constructor(private chargerService: ChargerService) {
    this.selectedCharger = {} as ChargerDTO
    this.chargers = []
    this.chargersToMap = []
  }

  ngOnInit(): void {
    this.loadChargers();
  }

  toggleSearch() {
    this.searchOpened = !this.searchOpened;
  }

  loadChargers() {
    this.chargerService.getAllChargers().subscribe(
      (response) => {
        this.chargers = response.body as ChargerDTO[]
        console.log(this.chargers)
        this.selectedCharger = this.chargers[0]
        this.chargers.forEach(chargerForMap => {
          this.chargersToMap.push(chargerForMap)
        });
        // set user location
        navigator.geolocation.getCurrentPosition((position) => {
          console.log(position.coords)
          // this.mapComponent.userLocation.push(position.coords.longitude)
          // this.mapComponent.userLocation.push(position.coords.latitude)
          this.mapComponent.userLocation.push(19.838230)
          this.mapComponent.userLocation.push(45.236300)
        })
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

  changeChargers(chargers: ChargerDTO[]) {
    this.mapComponent.chargers = chargers;
    this.mapComponent.showUserLocation = false;
    this.mapComponent.showClosestChargerLocation = false;
    this.mapComponent.refreshMapAfterSearch();
  }

  findMe() {
    this.mapComponent.showUserLocation = true;
    this.mapComponent.showClosestChargerLocation = false;
    this.mapComponent.findUserAndRefreshMap();
  }

  findClosestCharger() {
    this.findMe();
    this.chargerService.findClosestCharger(this.mapComponent.userLocation).subscribe(
      (response) => {
        let charger = response.body as ChargerDTO;
        this.chargersToMap.splice(this.chargersToMap.findIndex((element) => element.id === charger.id), 1);
        this.mapComponent.showUserLocation = true;
        this.mapComponent.showClosestChargerLocation = true;
        this.mapComponent.closestCharger = charger;
        this.mapComponent.refreshMapFeatures(this.chargersToMap);
        this.mapComponent.centerMapWithGivenCoordinates();
        // console.log(charger)
      },
      (err) => {
        console.log(err.error)
      }
    )
  }

  resetChargers() {
    this.mapComponent.chargers = this.chargers;
    this.mapComponent.showUserLocation = false;
    this.mapComponent.showClosestChargerLocation = false;
    this.mapComponent.refreshMapAfterSearch();
  }

}
