import { Component, Input, OnInit } from '@angular/core';
import { ChargerDTO } from 'src/modules/shared/model/chargerDTO';

@Component({
  selector: 'app-charger-info',
  templateUrl: './charger-info.component.html',
  styleUrls: ['./charger-info.component.scss']
})
export class ChargerInfoComponent implements OnInit {

  @Input()
  charger: ChargerDTO | undefined

  constructor() { }

  ngOnInit(): void {
  }

}
