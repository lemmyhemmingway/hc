import { Component } from "@angular/core";
import { HttpClientModule } from "@angular/common/http";
import { CommonModule } from "@angular/common";
import { UrlListComponent } from "./url-list/url-list.component";
import { RecordListComponent } from "./record-list/record-list.component";
import { UptimeTableComponent } from "./uptime-table/uptime-table.component";

@Component({
  selector: "app-root",
  standalone: true,
  imports: [
    CommonModule,
    HttpClientModule,
    UrlListComponent,
    RecordListComponent,
    UptimeTableComponent,
  ],
  templateUrl: "./app.component.html",
  styleUrls: ["./app.component.css"],
})
export class AppComponent {}
