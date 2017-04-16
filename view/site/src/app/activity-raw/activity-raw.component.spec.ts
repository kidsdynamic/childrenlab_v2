import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ActivityRawComponent } from './activity-raw.component';

describe('ActivityRawComponent', () => {
  let component: ActivityRawComponent;
  let fixture: ComponentFixture<ActivityRawComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ActivityRawComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ActivityRawComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
