import { Controller } from "@hotwired/stimulus";

export class TextFieldController extends Controller {
  static values = {
    inputErrorClass: { type: String, default: "input--error" },
  };
  static targets = ["input", "errorMessage"];

  declare inputErrorClassValue: string;
  declare hasErrorMessageTarget: boolean;
  declare errorMessageTarget: HTMLElement;
  declare inputTarget: HTMLInputElement;

  connect() {
    this.inputTarget.addEventListener("input", this.onInput);
  }

  disconnect() {
    this.inputTarget.removeEventListener("input", this.onInput);
  }

  onInput = () => {
    if (this.inputTarget.classList.contains(this.inputErrorClassValue)) {
      this.inputTarget.classList.remove(this.inputErrorClassValue);
    }

    if (this.hasErrorMessageTarget) {
      this.errorMessageTarget.classList.add("hidden");
    }
  };
}