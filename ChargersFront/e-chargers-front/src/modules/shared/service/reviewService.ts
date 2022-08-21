import { HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ReviewDTO } from '../model/reviewDTO';

@Injectable({
    providedIn: 'root'
})
export class ReviewService {

    private headers = new HttpHeaders({ "Content-Type": "application/json" });

    constructor(private http: HttpClient) { }

    create(reveiw: ReviewDTO): Observable<HttpResponse<string>> {
        let queryParams = {};

        queryParams = {
            headers: this.headers,
            observe: "response",
            responseType: "text"
        };

        return this.http.post<HttpResponse<string>>("echargers/api/recensions", reveiw, queryParams);
    }


}