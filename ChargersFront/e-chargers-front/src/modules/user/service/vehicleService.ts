import { HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ChargerDTO } from 'src/modules/shared/model/chargerDTO';
import { SearchDTO } from '../model/searchDTO';
import { VehicleDTO } from '../model/vehicleDTO';

@Injectable({
    providedIn: 'root'
})
export class VehicleService {

    private headers = new HttpHeaders({ "Content-Type": "application/json" });

    constructor(private http: HttpClient) { }

    getAllVehiclesOfUser(username: String): Observable<HttpResponse<VehicleDTO[]>> {
        let queryParams = {};

        queryParams = {
            headers: this.headers,
            observe: "response"
        };

        return this.http.get<HttpResponse<VehicleDTO[]>>("echargers/api/users/vehicles/" + username, queryParams);
    }

}