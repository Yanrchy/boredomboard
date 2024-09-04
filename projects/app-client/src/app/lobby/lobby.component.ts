import { Component } from '@angular/core';
import { ChannelService } from '../channel.service';

@Component({
  selector: 'app-lobby',
  standalone: true,
  imports: [],
  templateUrl: './lobby.component.html',
  styleUrl: './lobby.component.scss'
})
export class LobbyComponent {

  m_ChannelService: ChannelService;

  constructor(channelService: ChannelService) {

    this.m_ChannelService = channelService;

  }

}
