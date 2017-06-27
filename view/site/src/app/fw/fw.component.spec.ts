import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FWComponent } from './fw.component';

describe('FWComponent', () => {
  let component: FWComponent;
  let fixture: ComponentFixture<FWComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FWComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FWComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
