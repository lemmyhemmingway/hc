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
  template: `
    <h1>Health Check UI</h1>
    <h2>URLs</h2>
    <ul>
      <li *ngFor="let u of urls">{{u.target}}</li>
    </ul>
    <h2>Records</h2>
    <table>
      <thead>
        <tr><th>URL</th><th>Status</th><th>Timestamp</th></tr>
      </thead>
      <tbody>
        <tr *ngFor="let r of records">
          <td>{{r.url?.target}}</td>
          <td>{{r.status_code}}</td>
          <td>{{r.timestamp}}</td>
        </tr>
      </tbody>
    </table>
  `,
  styles: [
    'table { border-collapse: collapse; }',
    'th, td { border: 1px solid #ccc; padding: 4px 8px; }'
  ],
})
export class App implements OnInit {
  protected urls: URLItem[] = [];
  protected records: RecordItem[] = [];

  constructor(private http: HttpClient) {}

  ngOnInit() {
    this.http.get<URLItem[]>('/urls').subscribe(d => this.urls = d);
    this.http.get<RecordItem[]>('/records').subscribe(d => this.records = d);
  }
}
