import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApiService, URLItem } from './api.service';

@Component({
  selector: 'app-url-list',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './url-list.component.html',
  styleUrls: ['./url-list.component.css'],
})
export class UrlListComponent implements OnInit {
  urls: URLItem[] = [];

  constructor(private api: ApiService) {}

  ngOnInit() {
    this.api.getUrls().subscribe((d) => (this.urls = d));
  }
}
