import { ApplicationConfig, provideBrowserGlobalErrorListeners, provideZoneChangeDetection } from '@angular/core';
import { provideHttpClient } from '@angular/common/http';
import { provideRouter, Routes } from '@angular/router';
import { UptimeTableComponent } from './uptime-table/uptime-table.component';
import { RecordListComponent } from './record-list/record-list.component';

const routes: Routes = [
  { path: '', component: UptimeTableComponent },
  { path: 'latest', component: RecordListComponent },
];

export const appConfig: ApplicationConfig = {
  providers: [
    provideBrowserGlobalErrorListeners(),
    provideZoneChangeDetection({ eventCoalescing: true }),
    provideHttpClient(),
    provideRouter(routes),

  ]
};
