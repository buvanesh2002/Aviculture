import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserfirstpageComponent } from './userfirstpage.component';

describe('UserfirstpageComponent', () => {
  let component: UserfirstpageComponent;
  let fixture: ComponentFixture<UserfirstpageComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [UserfirstpageComponent]
    });
    fixture = TestBed.createComponent(UserfirstpageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
