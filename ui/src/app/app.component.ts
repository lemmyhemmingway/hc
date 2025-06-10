import { Component, OnInit } from "@angular/core";
import { HttpClient, HttpClientModule } from "@angular/common/http";
import { CommonModule } from "@angular/common";

interface EnvironmentItem {
  ID: number;
  Name: string;
}

interface URLItem {
  ID: number;
  Target: string;
  Environment?: EnvironmentItem;
}

interface RecordItem {
  ID: number;
  URLID: number;
  StatusCode: number;
  Timestamp: string;
  URL?: URLItem;
}

@Component({
  selector: "app-root",
  standalone: true,
  imports: [CommonModule, HttpClientModule],
  templateUrl: "./app.component.html",
  styleUrls: ["./app.component.css"],
})
export class AppComponent implements OnInit {
  urls: URLItem[] = [];
  records: RecordItem[] = [];

  constructor(private http: HttpClient) {}

  ngOnInit() {
    this.http.get<URLItem[]>("/urls").subscribe((d) => (this.urls = d));
    this.http
      .get<RecordItem[]>("/records")
      .subscribe((d) => (this.records = d));
  }
}
