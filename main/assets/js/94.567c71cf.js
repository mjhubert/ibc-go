(window.webpackJsonp=window.webpackJsonp||[]).push([[94],{657:function(e,t,a){"use strict";a.r(t);var s=a(1),n=Object(s.a)({},(function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[a("h1",{attrs:{id:"adr-011-ics-20-transfer-state-entry-for-total-amount-of-tokens-in-escrow"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#adr-011-ics-20-transfer-state-entry-for-total-amount-of-tokens-in-escrow"}},[e._v("#")]),e._v(" ADR 011: ICS-20 transfer state entry for total amount of tokens in escrow")]),e._v(" "),a("h2",{attrs:{id:"changelog"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#changelog"}},[e._v("#")]),e._v(" Changelog")]),e._v(" "),a("ul",[a("li",[e._v("2023-05-24: Initial draft")])]),e._v(" "),a("h2",{attrs:{id:"status"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#status"}},[e._v("#")]),e._v(" Status")]),e._v(" "),a("p",[e._v("Accepted and applied in v7.1 of ibc-go")]),e._v(" "),a("h2",{attrs:{id:"context"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#context"}},[e._v("#")]),e._v(" Context")]),e._v(" "),a("p",[e._v("Every ICS-20 transfer channel has its own escrow bank account. This account is used to lock tokens that are transferred out of a chain that acts as the source of the tokens (i.e. when the tokens being transferred have not returned to the originating chain). This design makes it easy to query the balance of the escrow accounts and find out the total amount of tokens in escrow in a particular channel. However, there are use cases where it would be useful to determine the total escrowed amount of a given denomination across all channels where those tokens have been transferred out.")]),e._v(" "),a("p",[e._v("For example: assuming that there are three channels between Cosmos Hub to Osmosis and 10 ATOM have been transferred from the Cosmos Hub to Osmosis on each of those channels, then we would like to know that 30 ATOM have been transferred (i.e. are locked in the escrow accounts of each channel) without needing to iterate over each escrow account to add up the balances of each.")]),e._v(" "),a("p",[e._v("For a sample use case where this feature would be useful, please refer to Osmosis' rate limiting use case described in "),a("a",{attrs:{href:"https://github.com/cosmos/ibc-go/issues/2664",target:"_blank",rel:"noopener noreferrer"}},[e._v("#2664"),a("OutboundLink")],1),e._v(".")]),e._v(" "),a("h2",{attrs:{id:"decision"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#decision"}},[e._v("#")]),e._v(" Decision")]),e._v(" "),a("h3",{attrs:{id:"state-entry-denom-amount"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#state-entry-denom-amount"}},[e._v("#")]),e._v(" State entry denom -> amount")]),e._v(" "),a("p",[e._v("The total amount of tokens in escrow (across all transfer channels) for a given denomination is stored in state in an entry keyed by the denomination: "),a("code",[e._v("totalEscrowForDenom/{denom}")]),e._v(".")]),e._v(" "),a("h3",{attrs:{id:"panic-if-amount-is-negative"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#panic-if-amount-is-negative"}},[e._v("#")]),e._v(" Panic if amount is negative")]),e._v(" "),a("p",[e._v("If a negative amount is ever attempted to be stored, then the keeper function will panic:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"go",base64:"aWYgY29pbi5BbW91bnQuSXNOZWdhdGl2ZSgpIHsKICBwYW5pYyhmbXQuU3ByaW50ZigmcXVvdDthbW91bnQgY2Fubm90IGJlIG5lZ2F0aXZlOiAlcyZxdW90OywgY29pbi5BbW91bnQpKQp9Cg=="}}),e._v(" "),a("h3",{attrs:{id:"delete-state-entry-if-amount-is-zero"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#delete-state-entry-if-amount-is-zero"}},[e._v("#")]),e._v(" Delete state entry if amount is zero")]),e._v(" "),a("p",[e._v("When setting the amount for a particular denomination, the value might be zero if all tokens that were transferred out of the chain have been transferred back. If this happens, then the state entry for this particular denomination will be deleted, since Cosmos SDK's "),a("code",[e._v("x/bank")]),e._v(" module prunes any non-zero balances:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"go",base64:"aWYgY29pbi5BbW91bnQuSXNaZXJvKCkgewogIHN0b3JlLkRlbGV0ZShrZXkpIC8vIGRlbGV0ZSB0aGUga2V5IHNpbmNlIENvc21vcyBTREsgeC9iYW5rIG1vZHVsZSB3aWxsIHBydW5lIGFueSBub24temVybyBiYWxhbmNlcwogIHJldHVybgp9Cg=="}}),e._v(" "),a("h3",{attrs:{id:"bundle-escrow-unescrow-with-setting-state-entry"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#bundle-escrow-unescrow-with-setting-state-entry"}},[e._v("#")]),e._v(" Bundle escrow/unescrow with setting state entry")]),e._v(" "),a("p",[e._v("Two new functions are implemented that bundle together the operations of escrowing/unescrowing and setting the total escrow amount in state, since these operations need to be executed together.")]),e._v(" "),a("p",[e._v("For escrowing tokens:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"go",base64:"Ly8gZXNjcm93VG9rZW4gd2lsbCBzZW5kIHRoZSBnaXZlbiB0b2tlbiBmcm9tIHRoZSBwcm92aWRlZCBzZW5kZXIgdG8gdGhlIGVzY3JvdyBhZGRyZXNzLiBJdCB3aWxsIGFsc28KLy8gdXBkYXRlIHRoZSB0b3RhbCBlc2Nyb3dlZCBhbW91bnQgYnkgYWRkaW5nIHRoZSBlc2Nyb3dlZCB0b2tlbiB0byB0aGUgY3VycmVudCB0b3RhbCBlc2Nyb3cuCmZ1bmMgKGsgS2VlcGVyKSBlc2Nyb3dUb2tlbihjdHggc2RrLkNvbnRleHQsIHNlbmRlciwgZXNjcm93QWRkcmVzcyBzZGsuQWNjQWRkcmVzcywgdG9rZW4gc2RrLkNvaW4pIGVycm9yIHsKICBpZiBlcnIgOj0gay5iYW5rS2VlcGVyLlNlbmRDb2lucyhjdHgsIHNlbmRlciwgZXNjcm93QWRkcmVzcywgc2RrLk5ld0NvaW5zKHRva2VuKSk7IGVyciAhPSBuaWwgewogICAgLy8gZmFpbHVyZSBpcyBleHBlY3RlZCBmb3IgaW5zdWZmaWNpZW50IGJhbGFuY2VzCiAgICByZXR1cm4gZXJyCiAgfQoKICAvLyB0cmFjayB0aGUgdG90YWwgYW1vdW50IGluIGVzY3JvdyBrZXllZCBieSBkZW5vbWluYXRpb24gdG8gYWxsb3cgZm9yIGVmZmljaWVudCBpdGVyYXRpb24KICBjdXJyZW50VG90YWxFc2Nyb3cgOj0gay5HZXRUb3RhbEVzY3Jvd0ZvckRlbm9tKGN0eCwgdG9rZW4uR2V0RGVub20oKSkKICBuZXdUb3RhbEVzY3JvdyA6PSBjdXJyZW50VG90YWxFc2Nyb3cuQWRkKHRva2VuKQogIGsuU2V0VG90YWxFc2Nyb3dGb3JEZW5vbShjdHgsIG5ld1RvdGFsRXNjcm93KQoKICByZXR1cm4gbmlsCn0K"}}),e._v(" "),a("p",[e._v("For unescrowing tokens:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"go",base64:"Ly8gdW5lc2Nyb3dUb2tlbiB3aWxsIHNlbmQgdGhlIGdpdmVuIHRva2VuIGZyb20gdGhlIGVzY3JvdyBhZGRyZXNzIHRvIHRoZSBwcm92aWRlZCByZWNlaXZlci4gSXQgd2lsbCBhbHNvCi8vIHVwZGF0ZSB0aGUgdG90YWwgZXNjcm93IGJ5IGRlZHVjdGluZyB0aGUgdW5lc2Nyb3dlZCB0b2tlbiBmcm9tIHRoZSBjdXJyZW50IHRvdGFsIGVzY3Jvdy4KZnVuYyAoayBLZWVwZXIpIHVuZXNjcm93VG9rZW4oY3R4IHNkay5Db250ZXh0LCBlc2Nyb3dBZGRyZXNzLCByZWNlaXZlciBzZGsuQWNjQWRkcmVzcywgdG9rZW4gc2RrLkNvaW4pIGVycm9yIHsKICBpZiBlcnIgOj0gay5iYW5rS2VlcGVyLlNlbmRDb2lucyhjdHgsIGVzY3Jvd0FkZHJlc3MsIHJlY2VpdmVyLCBzZGsuTmV3Q29pbnModG9rZW4pKTsgZXJyICE9IG5pbCB7CiAgICAvLyBOT1RFOiB0aGlzIGVycm9yIGlzIG9ubHkgZXhwZWN0ZWQgdG8gb2NjdXIgZ2l2ZW4gYW4gdW5leHBlY3RlZCBidWcgb3IgYSBtYWxpY2lvdXMKICAgIC8vIGNvdW50ZXJwYXJ0eSBtb2R1bGUuIFRoZSBidWcgbWF5IG9jY3VyIGluIGJhbmsgb3IgYW55IHBhcnQgb2YgdGhlIGNvZGUgdGhhdCBhbGxvd3MKICAgIC8vIHRoZSBlc2Nyb3cgYWRkcmVzcyB0byBiZSBkcmFpbmVkLiBBIG1hbGljaW91cyBjb3VudGVycGFydHkgbW9kdWxlIGNvdWxkIGRyYWluIHRoZQogICAgLy8gZXNjcm93IGFkZHJlc3MgYnkgYWxsb3dpbmcgbW9yZSB0b2tlbnMgdG8gYmUgc2VudCBiYWNrIHRoZW4gd2VyZSBlc2Nyb3dlZC4KICAgIHJldHVybiBlcnJvcnNtb2QuV3JhcChlcnIsICZxdW90O3VuYWJsZSB0byB1bmVzY3JvdyB0b2tlbnMsIHRoaXMgbWF5IGJlIGNhdXNlZCBieSBhIG1hbGljaW91cyBjb3VudGVycGFydHkgbW9kdWxlIG9yIGEgYnVnOiBwbGVhc2Ugb3BlbiBhbiBpc3N1ZSBvbiBjb3VudGVycGFydHkgbW9kdWxlJnF1b3Q7KQogIH0KCiAgLy8gdHJhY2sgdGhlIHRvdGFsIGFtb3VudCBpbiBlc2Nyb3cga2V5ZWQgYnkgZGVub21pbmF0aW9uIHRvIGFsbG93IGZvciBlZmZpY2llbnQgaXRlcmF0aW9uCiAgY3VycmVudFRvdGFsRXNjcm93IDo9IGsuR2V0VG90YWxFc2Nyb3dGb3JEZW5vbShjdHgsIHRva2VuLkdldERlbm9tKCkpCiAgbmV3VG90YWxFc2Nyb3cgOj0gY3VycmVudFRvdGFsRXNjcm93LlN1Yih0b2tlbikKICBrLlNldFRvdGFsRXNjcm93Rm9yRGVub20oY3R4LCBuZXdUb3RhbEVzY3JvdykKCiAgcmV0dXJuIG5pbAp9Cg=="}}),e._v(" "),a("p",[e._v("When tokens need to be escrowed in "),a("code",[e._v("sendTransfer")]),e._v(", then "),a("code",[e._v("escrowToken")]),e._v(" is called; when tokens need to be unescrowed on execution of the "),a("code",[e._v("OnRecvPacket")]),e._v(", "),a("code",[e._v("OnAcknowledgementPacket")]),e._v(" or "),a("code",[e._v("OnTimeoutPacket")]),e._v(" callbacks, then "),a("code",[e._v("unescrowToken")]),e._v(" is called.")]),e._v(" "),a("h3",{attrs:{id:"grpc-query-endpoint-and-cli-to-retrieve-amount"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#grpc-query-endpoint-and-cli-to-retrieve-amount"}},[e._v("#")]),e._v(" gRPC query endpoint and CLI to retrieve amount")]),e._v(" "),a("p",[e._v("A gRPC query endpoint is added so that it is possible to retrieve the total amount for a given denomination:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"proto",base64:"Ly8gVG90YWxFc2Nyb3dGb3JEZW5vbSByZXR1cm5zIHRoZSB0b3RhbCBhbW91bnQgb2YgdG9rZW5zIGluIGVzY3JvdyBiYXNlZCBvbiB0aGUgZGVub20uCnJwYyBUb3RhbEVzY3Jvd0ZvckRlbm9tKFF1ZXJ5VG90YWxFc2Nyb3dGb3JEZW5vbVJlcXVlc3QpIHJldHVybnMgKFF1ZXJ5VG90YWxFc2Nyb3dGb3JEZW5vbVJlc3BvbnNlKSB7CiAgb3B0aW9uIChnb29nbGUuYXBpLmh0dHApLmdldCA9ICZxdW90Oy9pYmMvYXBwcy90cmFuc2Zlci92MS9kZW5vbXMve2Rlbm9tPSoqfS90b3RhbF9lc2Nyb3cmcXVvdDs7Cn0KCi8vIFF1ZXJ5VG90YWxFc2Nyb3dGb3JEZW5vbVJlcXVlc3QgaXMgdGhlIHJlcXVlc3QgdHlwZSBmb3IgVG90YWxFc2Nyb3dGb3JEZW5vbSBSUEMgbWV0aG9kLgptZXNzYWdlIFF1ZXJ5VG90YWxFc2Nyb3dGb3JEZW5vbVJlcXVlc3QgewogIHN0cmluZyBkZW5vbSA9IDE7Cn0KCi8vIFF1ZXJ5VG90YWxFc2Nyb3dGb3JEZW5vbVJlc3BvbnNlIGlzIHRoZSByZXNwb25zZSB0eXBlIGZvciBUb3RhbEVzY3Jvd0ZvckRlbm9tIFJQQyBtZXRob2QuCm1lc3NhZ2UgUXVlcnlUb3RhbEVzY3Jvd0ZvckRlbm9tUmVzcG9uc2UgewogIGNvc21vcy5iYXNlLnYxYmV0YTEuQ29pbiBhbW91bnQgPSAxIFsoZ29nb3Byb3RvLm51bGxhYmxlKSA9IGZhbHNlXTsKfQo="}}),e._v(" "),a("p",[e._v("And a CLI query is also available to retrieve the total amount via the command line:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"shell",base64:"cXVlcnkgaWJjLXRyYW5zZmVyIHRvdGFsLWVzY3JvdyBbZGVub21dCg=="}}),e._v(" "),a("h2",{attrs:{id:"consequences"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#consequences"}},[e._v("#")]),e._v(" Consequences")]),e._v(" "),a("h3",{attrs:{id:"positive"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#positive"}},[e._v("#")]),e._v(" Positive")]),e._v(" "),a("ul",[a("li",[e._v("Possibility to retrieve the total amount of a particular denomination in escrow across all transfer channels without iteration.")])]),e._v(" "),a("h3",{attrs:{id:"negative"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#negative"}},[e._v("#")]),e._v(" Negative")]),e._v(" "),a("p",[e._v("No notable consequences")]),e._v(" "),a("h3",{attrs:{id:"neutral"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#neutral"}},[e._v("#")]),e._v(" Neutral")]),e._v(" "),a("ul",[a("li",[e._v("A new entry is added to state for every denomination that is transferred out of the chain.")])]),e._v(" "),a("h2",{attrs:{id:"references"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#references"}},[e._v("#")]),e._v(" References")]),e._v(" "),a("p",[e._v("Issues:")]),e._v(" "),a("ul",[a("li",[a("a",{attrs:{href:"https://github.com/cosmos/ibc-go/issues/2664",target:"_blank",rel:"noopener noreferrer"}},[e._v("#2664"),a("OutboundLink")],1)])]),e._v(" "),a("p",[e._v("PRs:")]),e._v(" "),a("ul",[a("li",[a("a",{attrs:{href:"https://github.com/cosmos/ibc-go/pull/3019",target:"_blank",rel:"noopener noreferrer"}},[e._v("#3019"),a("OutboundLink")],1)]),e._v(" "),a("li",[a("a",{attrs:{href:"https://github.com/cosmos/ibc-go/pull/3558",target:"_blank",rel:"noopener noreferrer"}},[e._v("#3558"),a("OutboundLink")],1)])])],1)}),[],!1,null,null,null);t.default=n.exports}}]);