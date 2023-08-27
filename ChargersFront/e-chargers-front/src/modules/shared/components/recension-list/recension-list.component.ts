import { Component, Input, OnInit } from '@angular/core';
import { RecensionWithUserDTO } from '../../model/reviewDTO';

@Component({
  selector: 'app-recension-list',
  templateUrl: './recension-list.component.html',
  styleUrls: ['./recension-list.component.scss']
})
export class RecensionListComponent implements OnInit {

  @Input()
  recensions: RecensionWithUserDTO[] = [];

  constructor() { }

  ngOnInit(): void {
  }

}
