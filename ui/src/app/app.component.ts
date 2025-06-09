import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { CommonModule } from '@angular/common';

interface URLItem {
  id: number;
  target: string;
}

interface RecordItem {
  id: number;
  url_id: number;
  status_code: number;
  timestamp: string;
  url?: URLItem;
}

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, HttpClientModule],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  urls: URLItem[] = [];
  records: RecordItem[] = [];

  constructor(private http: HttpClient) {}

  ngOnInit() {
    this.http.get<URLItem[]>('/urls').subscribe(d => this.urls = d);
    this.http.get<RecordItem[]>('/records').subscribe(d => this.records = d);
  }
}
