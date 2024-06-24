import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class UserProfileService {

  m_Username: string;
  m_Color:    string;

  constructor() {

    this.m_Username = "";
    this.m_Color    = "";

   }

   Register(username: string, color: string): void {

    this.m_Username = username;
    this.m_Color    = color;

   }

}
