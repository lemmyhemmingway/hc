import { Component, OnInit } from "@angular/core";
import { CommonModule } from "@angular/common";
import { ApiService, UptimeItem } from "../api.service";

@Component({
  selector: "app-uptime-table",
  standalone: true,
  imports: [CommonModule],
  templateUrl: "./uptime-table.component.html",
  styleUrls: ["./uptime-table.component.css"],
})
export class UptimeTableComponent implements OnInit {
  items: UptimeItem[] = [];

  constructor(private api: ApiService) {}

  ngOnInit() {
    this.api.getUptimes().subscribe((d) => (this.items = d));
  }
}
