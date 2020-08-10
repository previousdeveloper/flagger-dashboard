import {Component, OnInit} from '@angular/core';
import {CanaryService} from "../dashboard/canary.service";

@Component({
    selector: 'app-tables',
    templateUrl: './tables.component.html',
    styleUrls: ['./tables.component.scss']
})
export class TablesComponent implements OnInit {
    canaryDeployments: any[] = [];
    canaryStatusBagde: {} = {
        "Progressing": "bg-primary",
        "Promoting": "bg-info",
        "Succeeded": "bg-success",
        "Failed": "bg-danger",
    };

    private http: CanaryService;

    constructor(http: CanaryService) {
        this.http = http;
        this.fetchDataPeriodically()

    }

    ngOnInit() {
        this.http.ngOnInit().subscribe(x => {
            this.canaryDeployments = this.getCanaryDeployments(x)
        });
    }

    getCanaryDeployments(results: any) {
        let canaryDeployments: any[] = [];

        for (let entry of results) {
            let each = {
                "name": entry.metadata.name,
                "namespace": entry.metadata.namespace,
                "badge": this.canaryStatusBagde[entry.status.conditions[0].reason],
                "status": entry.status.conditions[0].reason,
                "lastTransitionTime": entry.status.conditions[0].lastTransitionTime,
                "message": entry.status.conditions[0].message,
                "canaryWeight": entry.status.canaryWeight,
                "failedChecks": entry.status.failedChecks
            };
            canaryDeployments.push(each);

        }

        return canaryDeployments;
    }

    fetchDataPeriodically() {

        setInterval(() => {
            this.http.ngOnInit().subscribe(x => {
                this.canaryDeployments = this.getCanaryDeployments(x)
            });
        }, 5000);

    }

}
