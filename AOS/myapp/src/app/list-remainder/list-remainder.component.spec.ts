import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListRemainderComponent } from './list-remainder.component';

describe('ListRemainderComponent', () => {
  let component: ListRemainderComponent;
  let fixture: ComponentFixture<ListRemainderComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ListRemainderComponent]
    });
    fixture = TestBed.createComponent(ListRemainderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
