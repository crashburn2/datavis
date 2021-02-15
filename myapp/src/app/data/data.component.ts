import { Component, OnInit } from '@angular/core';
import { HeroService } from '../hero.service';

@Component({
  selector: 'app-data',
  templateUrl: './data.component.html',
  styleUrls: ['./data.component.scss']
})
export class DataComponent implements OnInit {
  constructor(
    private heroService: HeroService) { 
      this.subscribeToData();
    }

  data: string[] | undefined;

  subscribeToData(): void {
    const dataObservable = this.heroService.getData();
    dataObservable.subscribe(data => this.data = data)

  }

  async getData(): Promise<string[] | undefined> {
    var privateData: string[] | undefined

    return new Promise(resolve => {
      const dataObservable = this.heroService.getData();
      dataObservable.subscribe(data => resolve(data))});
  }

  async ngOnInit(): Promise<void> {
    await this.getData();
    if (!this.data) {
      console.log("Fehler: data Empty")
      return
    }  }

}
