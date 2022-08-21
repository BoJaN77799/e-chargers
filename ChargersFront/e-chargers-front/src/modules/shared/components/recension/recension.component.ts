import { Component, Input, OnInit } from '@angular/core';
import * as moment from 'moment';
import { ReviewDTO } from '../../model/reviewDTO';

@Component({
  selector: 'app-recension',
  templateUrl: './recension.component.html',
  styleUrls: ['./recension.component.scss']
})
export class RecensionComponent implements OnInit {

  @Input()
  recension: ReviewDTO | undefined;

  constructor() { }

  ngOnInit(): void {
  }

  formatDate(date: number | undefined) {
    return moment(date).format("YYYY-MM-DD HH:mm")
  }

}
