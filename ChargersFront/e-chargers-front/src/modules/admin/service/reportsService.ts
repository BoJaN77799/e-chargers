import { HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Report } from '../model/reportsDTO';

@Injectable({
    providedIn: 'root'
})
export class ReportsService {

    private headers = new HttpHeaders({ "Content-Type": "application/json" });

    constructor(private http: HttpClient) { }


    getReport(date_from: number, date_to: number): Observable<HttpResponse<Report>> {
        let queryParams = {};

        queryParams = {
            headers: this.headers,
            observe: "response",
        };

        return this.http.get<HttpResponse<Report>>("echargers/api/reports/" + date_from + "/" + date_to, queryParams);
    }
}