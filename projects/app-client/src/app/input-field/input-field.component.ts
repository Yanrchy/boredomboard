import { Component, ElementRef, ViewChild } from '@angular/core';

import { ChannelService } from '../channel.service';

@Component({
  selector: 'app-input-field',
  standalone: true,
  imports: [],
  templateUrl: './input-field.component.html',
  styleUrl: './input-field.component.scss'
})
export class InputFieldComponent {

  m_ChannelService: ChannelService;
  @ViewChild("inputText") m_ElementRef!: ElementRef;

  constructor(channelService: ChannelService) {

    this.m_ChannelService = channelService;

  }

  Send(message: string) {

    this.m_ChannelService.Send(message);
    
    this.m_ElementRef.nativeElement.value = "";

  }

}
