import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import 'ol/ol.css';
import Map from 'ol/Map';
import View from 'ol/View';
import { OSM } from 'ol/source';
import Tile from 'ol/layer/Tile';
import { fromLonLat, transform } from 'ol/proj';
import { Style, Icon } from 'ol/style'
import { Feature } from 'ol';
import { Point } from 'ol/geom'
import VectorSource from 'ol/source/Vector';
import VectorLayer from 'ol/layer/Vector';
import { ChargerDTO } from '../../model/chargerDTO';
import { ChargerService } from 'src/modules/user/service/chargerService';
import { INPUT_MODALITY_DETECTOR_DEFAULT_OPTIONS } from '@angular/cdk/a11y';

@Component({
  selector: 'app-map-page',
  templateUrl: './map-page.component.html',
  styleUrls: ['./map-page.component.scss']
})
export class MapPageComponent implements OnInit {

  @Input()
  chargers: ChargerDTO[] = [];

  userLocation: number[] = [];

  showUserLocation = false
  showClosestChargerLocation = false

  closestCharger!: ChargerDTO;

  @Output()
  selectedChargerEvent = new EventEmitter<ChargerDTO>();

  public map!: Map

  constructor(private chargerService: ChargerService) { }

  ngOnInit(): void {
    this.loadChargers()
  }

  loadChargers() {
    this.chargerService.getAllChargers().subscribe(
      (response) => {
        this.chargers = response.body as ChargerDTO[]
        this.initMap(this.chargers)
      },
      (err) => {
        console.log(err.error)
      }
    )
  }

  crateFeaturesFromChargers(chargers: ChargerDTO[]) {
    var featureList = [];

    for (let idx = 0; idx < chargers.length; idx++) {
      const charger = chargers[idx];

      let feature = new Feature({
        geometry: new Point(fromLonLat([charger.address.longitude, charger.address.latitude]))
      })
      feature.setStyle(new Style({
        image: new Icon(({
          anchor: [0.5, 1],
          src: 'assets/img/marker-green-smalle.png',
        }))
      }))

      feature.setProperties({ 'name': charger.name })

      featureList.push(feature)
    }
    return featureList
  }

  createUserLocationFeature(userLocation: number[]) {
    let userFeature = new Feature({
      geometry: new Point(fromLonLat([userLocation[0], userLocation[1]]))
    })
    userFeature.setStyle(new Style({
      image: new Icon(({
        anchor: [0.5, 1],
        src: 'assets/img/red-marker-small.png',
      }))
    }))
    return userFeature;
  }

  createClosestChargerFeature(userLocation: number[]) {
    let feature = new Feature({
      geometry: new Point(fromLonLat([userLocation[0], userLocation[1]]))
    })
    feature.setStyle(new Style({
      image: new Icon(({
        anchor: [0.5, 1],
        src: 'assets/img/red-marker-small.png',
      }))
    }))
    return feature;
  }

  initMap(chargers: ChargerDTO[]) {

    var featureList = this.crateFeaturesFromChargers(chargers);

    this.map = new Map({
      layers: [
        new Tile({
          source: new OSM(),
        }),
        new VectorLayer({
          source: new VectorSource({
            features: [...featureList]
          })
        })],
      target: 'map',
      view: new View({
        center: fromLonLat([19.845820, 45.244630]),
        zoom: 13
      }),
    });

    this.map.on('click', e => {
      this.map.forEachFeatureAtPixel(e.pixel,
        (feature, layer) => {
          // console.log(feature.get('name'))
          let selectedCharger = this.findChargerByName(feature.get('name'), chargers)
          if (selectedCharger)
            this.selectedChargerEvent.emit(selectedCharger)
        })
    })

    this.map.on("pointermove", evt => {
      var hit = this.map.forEachFeatureAtPixel(evt.pixel, function (feature, layer) {
        return true;
      });
      if (hit) {
        this.map.getTargetElement().style.cursor = 'pointer';
      } else {
        this.map.getTargetElement().style.cursor = '';
      }
    });

  }

  findChargerByName(name: string, chargers: ChargerDTO[]): ChargerDTO | undefined {
    for (let idx = 0; idx < chargers.length; idx++) {
      const element = chargers[idx];
      if (element.name == name)
        return element
    }
    return undefined
  }

  refreshMapFeatures(chargers: ChargerDTO[]) {
    var featureList = this.crateFeaturesFromChargers(chargers)

    if (this.showUserLocation) {
      featureList.push(this.createUserLocationFeature(this.userLocation))
    }

    if (this.showClosestChargerLocation) {
      let lon = this.closestCharger.address.longitude
      let lat = this.closestCharger.address.latitude
      featureList.push(this.createClosestChargerFeature([lon, lat]))
    }

    this.map.getLayers()['array_'][1] = new VectorLayer({
      source: new VectorSource({
        features: [...featureList]
      })
    })

    this.map.updateSize();
  }

  centerMap() {
    this.map.getView().setCenter(transform([this.userLocation[0], this.userLocation[1]], 'EPSG:4326', 'EPSG:3857'));
    this.map.getView().setZoom(15);
  }

  centerMapWithGivenCoordinates() {
    if (this.closestCharger) {
      let centerLon = (this.userLocation[0] + this.closestCharger?.address.longitude) / 2
      let centerLat = (this.userLocation[1] + this.closestCharger?.address.latitude) / 2
      this.map.getView().setCenter(transform([centerLon, centerLat], 'EPSG:4326', 'EPSG:3857'));
      this.map.getView().setZoom(15);
    }
  }


  refreshMapAfterSearch() {
    this.refreshMapFeatures(this.chargers)
  }

  findUserAndRefreshMap() {
    this.refreshMapFeatures(this.chargers)
    this.centerMap()
  }

}

