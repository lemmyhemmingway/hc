import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';

interface UptimeItem {
  up: boolean;
  environment: string;
  tags: string[];
  location: string;
  type: string;
  uptime: string;
  upSince: string;
  responseTime: string;
}

@Component({
  selector: 'app-uptime-table',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './uptime-table.component.html',
  styleUrls: ['./uptime-table.component.css'],
})
export class UptimeTableComponent {
  items: UptimeItem[] = [
    {
      up: true,
      environment: 'Prod',
      tags: ['api', 'critical'],
      location: 'US-East',
      type: 'https',
      uptime: '99.9%',
      upSince: '2025-06-01',
      responseTime: '250ms',
    },
    {
      up: false,
      environment: 'Staging',
      tags: ['internal'],
      location: 'EU-West',
      type: 'ping',
      uptime: '97.5%',
      upSince: '2025-05-27',
      responseTime: '1.2s',
    },
  ];
}
