import { HttpBackend, HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
import { add } from 'ol/coordinate';
import { Observable } from 'rxjs';
import { Address } from '../model/chargerDTO';

@Injectable({
    providedIn: 'root'
})
export class UtilService {

    private headers = new HttpHeaders({ "Content-Type": "application/json" });

    constructor(private http: HttpClient, private handler: HttpBackend) {
        this.http = new HttpClient(handler)
    }

    public getNoPages(totalItems: number, pageSize: number): number {
        return Math.ceil(totalItems / pageSize);
    }

    public isRole(role: string): boolean {
        const item = sessionStorage.getItem("user");

        if (item) {
            const jwt: JwtHelperService = new JwtHelperService();
            const roleFromToken: string = jwt.decodeToken(JSON.parse(item).token).role;
            return role === roleFromToken
        }
        return false;
    }

    public getLoggedUsername(): string {
        const item = sessionStorage.getItem("user");
        if (item) {
            const jwt: JwtHelperService = new JwtHelperService();
            const userToken = JSON.parse(item);
            return jwt.decodeToken(userToken['token']).username;
        }
        return "";
    }

    isLoggedIn(): boolean {
        if (!sessionStorage.getItem("user")) {
            return false;
        }
        return true;
    }

    findCoordinated(address: Address): Observable<HttpResponse<any>> {

        let queryParams = {};

        queryParams = {
            headers: this.headers,
            observe: "response",
            format: "json",
            limit: 1,
            "accept-language": "en",
            "Access-Control-Allow-Origin": "*"
        };

        const url =
            "https://nominatim.openstreetmap.org/search?q=" +
            address.street +
            ", " +
            address.city + "&format=json&polygon=1&addressdetails=1"

        return this.http.get<HttpResponse<any>>(url, queryParams);

    }
}
