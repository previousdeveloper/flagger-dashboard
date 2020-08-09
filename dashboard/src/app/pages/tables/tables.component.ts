import { Component, OnInit } from '@angular/core';
import {CanaryService} from "../dashboard/canary.service";

@Component({
  selector: 'app-tables',
  templateUrl: './tables.component.html',
  styleUrls: ['./tables.component.scss']
})
export class TablesComponent implements OnInit {
  data2: string;

  private http: CanaryService;

  constructor(http: CanaryService) {
    this.http = http;
  }

  ngOnInit() {
    this.http.ngOnInit().subscribe(x => this.data2 = x);

  }

}
