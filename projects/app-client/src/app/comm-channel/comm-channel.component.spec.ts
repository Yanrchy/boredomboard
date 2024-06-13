import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CommChannelComponent } from './comm-channel.component';

describe('CommChannelComponent', () => {
  let component: CommChannelComponent;
  let fixture: ComponentFixture<CommChannelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CommChannelComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(CommChannelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
