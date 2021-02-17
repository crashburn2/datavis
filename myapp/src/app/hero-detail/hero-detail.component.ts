import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Hero } from '../hero';
import { HeroService } from '../hero.service';

@Component({
  selector: 'app-hero-detail',
  templateUrl: './hero-detail.component.html',
  styleUrls: ['./hero-detail.component.scss']
})
export class HeroDetailComponent implements OnInit {

  hero: Hero | undefined;
  heroId: string | undefined

  constructor(
    private route: ActivatedRoute,
    private heroService: HeroService,
    private location: Location
  ) { }

  async ngOnInit(): Promise<void> {
    const id: string | null = this.route.snapshot.paramMap.get('id');
    if (!id) {
      return
    }
    this.heroId = id
    this.hero = await this.heroService.getHero(+id)
  }

  goBack(): void {
    this.location.back();
  }

  async save(): Promise<void> {
    if (!this.hero) {
      return
    }

    await this.heroService.putHero(this.hero)
  }
}