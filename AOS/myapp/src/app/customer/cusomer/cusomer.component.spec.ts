import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CusomerComponent } from './cusomer.component';

describe('CusomerComponent', () => {
  let component: CusomerComponent;
  let fixture: ComponentFixture<CusomerComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CusomerComponent]
    });
    fixture = TestBed.createComponent(CusomerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
