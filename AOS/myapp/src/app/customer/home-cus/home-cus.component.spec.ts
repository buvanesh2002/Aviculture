import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomeCusComponent } from './home-cus.component';

describe('HomeCusComponent', () => {
  let component: HomeCusComponent;
  let fixture: ComponentFixture<HomeCusComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [HomeCusComponent]
    });
    fixture = TestBed.createComponent(HomeCusComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
