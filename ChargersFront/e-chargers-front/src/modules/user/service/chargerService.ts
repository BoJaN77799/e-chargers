import { HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Charger, ChargerDTO } from 'src/modules/shared/model/chargerDTO';
import { SearchDTO } from '../model/searchDTO';

@Injectable({
    providedIn: 'root'
})
export class ChargerService {

    private headers = new HttpHeaders({ "Content-Type": "application/json" });

    constructor(private http: HttpClient) { }

    getAllChargers(): Observable<HttpResponse<ChargerDTO[]>> {
        let queryParams = {};

        queryParams = {
            headers: this.headers,
            observe: "response"
        };

        return this.http.get<HttpResponse<ChargerDTO[]>>("echargers/api/chargers", queryParams);
    }

    search(searchDTO: SearchDTO): Observable<HttpResponse<ChargerDTO[]>> {
        let queryParams = {};

        queryParams = {
            headers: this.headers,
            observe: "response"
        };

        return this.http.post<HttpResponse<ChargerDTO[]>>("echargers/api/chargers/search", searchDTO, queryParams);
    }

    getChargerById(charger_id: number): Observable<HttpResponse<ChargerDTO>> {
        let queryParams = {};

        queryParams = {
            headers: this.headers,
            observe: "response"
        };

        return this.http.get<HttpResponse<ChargerDTO>>("echargers/api/chargers/" + charger_id, queryParams);
    }

    create(charger: Charger): Observable<HttpResponse<string>> {
        let queryParams = {};

        queryParams = {
            headers: this.headers,
            observe: "response",
            responseType: "text"
        };

        return this.http.post<HttpResponse<string>>("echargers/api/chargers", charger, queryParams);
    }


}