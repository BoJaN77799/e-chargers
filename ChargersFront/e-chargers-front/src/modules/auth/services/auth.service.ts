import { Injectable } from "@angular/core";
import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Observable } from "rxjs";
import { Login } from "../models/login";
import { Token } from "../models/token";
import { RegistrationDTO } from "../models/registration";

@Injectable({
  providedIn: "root",
})
export class AuthService {
  private headers = new HttpHeaders({ "Content-Type": "application/json" });

  constructor(private http: HttpClient) { }

  login(auth: Login): Observable<Token> {
    return this.http.post<Token>("echargers/api/users/login", auth, {
      headers: this.headers,
      responseType: "json",
    });
  }

  register(RegistrationDTO: RegistrationDTO): Observable<HttpResponse<string>> {
    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: "response",
      responseType: "text"
    };

    return this.http.post<HttpResponse<string>>("echargers/api/users", RegistrationDTO, queryParams);
  }

  logout(): Observable<String> {
    return this.http.post<String>("echargers/api/users/logout", {
      headers: this.headers,
      responseType: "json",
    });
  }
}
