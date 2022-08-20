import { HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';

@Injectable({
    providedIn: 'root'
})
export class UtilService {

    constructor(private http: HttpClient) { }

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

}
