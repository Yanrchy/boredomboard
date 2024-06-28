import { Component } from '@angular/core';

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

  constructor(channelService: ChannelService) {

    this.m_ChannelService = channelService;

  }

  Send(message: string) {

    this.m_ChannelService.Send(message);

  }

}
