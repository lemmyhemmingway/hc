import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApiService, RecordItem } from '../api.service';

@Component({
  selector: 'app-record-list',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './record-list.component.html',
  styleUrls: ['./record-list.component.css'],
})
export class RecordListComponent implements OnInit {
  records: RecordItem[] = [];

  constructor(private api: ApiService) {}

  ngOnInit() {
    this.api.getRecords().subscribe((d) => (this.records = d));
  }
}
