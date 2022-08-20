import { HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ChargerDTO } from 'src/modules/shared/model/chargerDTO';
import { ReservationDTO } from '../model/reservationDTO';

@Injectable({
    providedIn: 'root'
})
export class ReservationService {

    private headers = new HttpHeaders({ "Content-Type": "application/json" });

    constructor(private http: HttpClient) { }

    createReservation(reservationDTO: ReservationDTO): Observable<HttpResponse<string>> {
        let queryParams = {};

        queryParams = {
            headers: this.headers,
            observe: "response",
            responseType: "text"
        };

        return this.http.post<HttpResponse<string>>("echargers/api/reservations", reservationDTO, queryParams);
    }

    getAllReservationsOfUser(username: string): Observable<HttpResponse<ReservationDTO[]>> {
        let queryParams = {};

        queryParams = {
            headers: this.headers,
            observe: "response",
        };

        return this.http.get<HttpResponse<ReservationDTO[]>>("echargers/api/reservations/" + username, queryParams);
    }

    cancelReservation(id: number): Observable<HttpResponse<string>> {
        let queryParams = {};

        queryParams = {
            headers: this.headers,
            observe: "response",
            responseType: "text"
        };

        return this.http.delete<HttpResponse<string>>("echargers/api/reservations/" + id, queryParams);
    }
}