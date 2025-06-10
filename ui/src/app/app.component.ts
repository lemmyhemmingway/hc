import { Component } from "@angular/core";
import { HttpClientModule } from "@angular/common/http";
import { CommonModule } from "@angular/common";
import { UrlListComponent } from "./url-list.component";
import { RecordListComponent } from "./record-list.component";

@Component({
  selector: "app-root",
  standalone: true,
  imports: [CommonModule, HttpClientModule, UrlListComponent, RecordListComponent],
  templateUrl: "./app.component.html",
  styleUrls: ["./app.component.css"],
})
export class AppComponent {}
