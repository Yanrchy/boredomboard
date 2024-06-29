import { Injectable } from '@angular/core';
import { UserProfileService } from './user-profile.service';

@Injectable({
  providedIn: 'root'
})
export class ChannelService {

  m_ProfileService: UserProfileService;
  m_Messages: Message[];

  constructor(profileService: UserProfileService) {

    this.m_ProfileService = profileService;
    this.m_Messages = new Array<Message>();

   }

   Send(text: string): void {

    this.m_Messages.push(new Message(this.m_ProfileService.m_Username, this.m_ProfileService.m_Color, text));

   }

}

class Message {

  m_Sender: string;
  m_Color:  string;
  m_Text:   string;

  constructor(sender: string, color: string, text: string) {

    this.m_Sender = sender;
    this.m_Color = color;
    this.m_Text = text;

  }

}
