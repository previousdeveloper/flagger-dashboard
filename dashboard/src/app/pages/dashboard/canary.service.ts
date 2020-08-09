import {Injectable, OnInit} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from "rxjs";

@Injectable({providedIn: 'root'})
export class CanaryService {
    updatedItem: string;

    constructor(private http: HttpClient) {
    }

    ngOnInit(): Observable<any> {
        return this.http
            .get<any>("http://localhost:8888/canary/browsing-team");
    }
}
