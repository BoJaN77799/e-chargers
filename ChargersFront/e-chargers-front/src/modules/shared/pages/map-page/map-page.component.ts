import { AfterViewInit, Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import 'ol/ol.css';
import Map from 'ol/Map';
import View from 'ol/View';
import { OSM } from 'ol/source';
import Tile from 'ol/layer/Tile';
import { fromLonLat, transform } from 'ol/proj';
import { Style, Icon, Stroke } from 'ol/style'
import { Feature, Overlay } from 'ol';
import { Point } from 'ol/geom'
import VectorSource from 'ol/source/Vector';
import VectorLayer from 'ol/layer/Vector';
import { Coordinates } from '../../model/coordinates';
import { ChargerDTO } from '../../model/chargerDTO';
import { ChargerService } from 'src/modules/user/service/chargerService';
import { map } from 'rxjs';

@Component({
  selector: 'app-map-page',
  templateUrl: './map-page.component.html',
  styleUrls: ['./map-page.component.scss']
})
export class MapPageComponent implements OnInit {

  @Input()
  chargers: ChargerDTO[] = [];

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
        console.log(response.body)
        this.chargers = response.body as ChargerDTO[]
        this.initMap()
      },
      (err) => {
        console.log(err.error)
      }
    )
  }

  initMap() {

    var featureList = [];

    for (let idx = 0; idx < this.chargers.length; idx++) {
      const charger = this.chargers[idx];

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
        zoom: 15
      }),
    });

    this.map.on('click', e => {
      this.map.forEachFeatureAtPixel(e.pixel,
        (feature, layer) => {
          console.log(feature.get('name'))
          let selectedCharger = this.findChargerByName(feature.get('name'))
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

  CenterMap(long: number, lat: number) {
    console.log("Long: " + long + " Lat: " + lat);
    this.map.getView().setCenter(transform([long, lat], 'EPSG:4326', 'EPSG:3857'));
    this.map.getView().setZoom(15);
  }

  findChargerByName(name: string): ChargerDTO | undefined {
    for (let idx = 0; idx < this.chargers.length; idx++) {
      const element = this.chargers[idx];
      if (element.name == name)
        return element
    }
    return undefined
  }

}

