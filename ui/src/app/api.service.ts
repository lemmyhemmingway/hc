import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Observable } from "rxjs";

export interface EnvironmentItem {
  ID: number;
  Name: string;
}

export interface URLItem {
  ID: number;
  Target: string;
  Environment?: EnvironmentItem;
}

export interface RecordItem {
  ID: number;
  URLID: number;
  StatusCode: number;
  Timestamp: string;
  URL?: URLItem;
}

export interface UptimeItem {
  up: boolean;
  environment: string;
  tags: string[];
  location: string;
  type: string;
  uptime: string;
  upSince: string;
  responseTime: string;
}

@Injectable({ providedIn: "root" })
export class ApiService {
  constructor(private http: HttpClient) {}

  getUrls(): Observable<URLItem[]> {
    return this.http.get<URLItem[]>("/urls");
  }

  getRecords(): Observable<RecordItem[]> {
    return this.http.get<RecordItem[]>("/records");
  }

  getUptimes(): Observable<UptimeItem[]> {
    return this.http.get<UptimeItem[]>("/uptimes");
  }
}
