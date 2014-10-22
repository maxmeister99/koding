!function(){var a,b=Array.prototype.slice,c=Object.prototype.hasOwnProperty;a=function(){function f(a,b,c){var d;this.view=a,this.template=b,this.options=null!=c?c:{},d=this.options,this.prefix=d.prefix,this.params=d.params,this.params||(this.params={}),this.symbols={},this.symbolsByDataPath={},this.symbolsBySubViewName={},this.dataPaths={},this.subViewNames={},this.prefix||(this.prefix=""),this.html=this.init()}var a,d,e;return f.createId=function(){var a;return a=0,function(b){return""+b+"el-"+a++}}(),f.getAt=function(a,b){var c;for(b="function"==typeof b.split?b.split("."):b.slice();null!=a&&(c=b.shift());)a=a[c];return a},d=/\{([\w|-]*)?(\#[\w|-]*)?((?:\.[\w|-]*)*)(\[(?:\b[\w|-]*\b)(?:\=[\"|\']?.*[\"|\']?)\])*\{([^{}]*)\}\s*\}/g,f.prototype.createId=f.createId,f.prototype.toString=function(){return this.template},f.prototype.init=function(){var c,e,g;return c=function(a){var b;return b="function"==typeof this.getData?this.getData():void 0,null!=b?("function"==typeof b.getAt?b.getAt(a):void 0)||f.getAt(b,a):void 0},e=function(a,b,c,d){return function(e){return b.embedChild(c,e,d.isCustom),d.isCustom?void 0:(d.id=e.id,d.tagName="function"==typeof e.getTagName?e.getTagName():void 0,delete a.symbols[c],a.symbols[e.id]=d)}},g=function(){var f,g,h,i=this;return g=this.prefix,h=this.view,f=this.createId,this.template.replace(d,function(d,j,k,l,m,n){var o,p,q,r,s,t,u,v,w,x,y,z,A,B,C;k=null!=k?k.split("#")[1]:void 0,p=(null!=l?l.split(".").slice(1):void 0)||[],m=(null!=m?m.replace(/\]\[/g," ").replace(/\[|\]/g,""):void 0)||"",v=!!(j||k||l.length||m.length),j||(j="span"),r=[],A=[],n=n.replace(/#\(([^)]*)\)/g,function(a,b){return r.push(b),"data('"+b+"')"}).replace(/^(?:> ?|embedChild )(.+)/,function(a,b){return A.push(b.replace(/\@\.?|this\./,"")),"embedChild("+b+")"}),i.registerDataPaths(r),i.registerSubViewNames(A),w="return "+n,"debug"===j&&(console.debug(w),j="span"),x=Object.keys(i.params),y=x.map(function(a){return i.params[a]}),u=["data","embedChild"].concat(b.call(x));try{q=Function.apply(null,b.call(u).concat([w]))}catch(D){throw new Error("Pistachio encountered an error: "+D+"\nSource: "+w)}return k||(k=f(g)),z=function(){return""+q.apply(h,[c.bind(h),t].concat(b.call(y)))},C={tagName:j,id:k,isCustom:v,js:w,code:q,render:z,dataPaths:r,subViewNames:A},i.addSymbolInternal(C),t=e(i,h,k,C),s=r.length?" data-"+g+"paths='"+r.join(" ")+"'":"",B=A.length?(p.push(""+g+"subview")," data-"+g+"subviews='"+a(A.join(" "))+"'"):"",o=p.length?" class='"+p.join(" ")+"'":"","<"+j+o+s+B+" "+m+" id='"+k+"'></"+j+">"})}}(),f.prototype.addSymbolInternal=function(a){var b,c,d,e,f,g,h,i,j,k;for(c=a.dataPaths,e=a.subViewNames,this.symbols[a.id]=a,h=0,j=c.length;j>h;h++)b=c[h],null==(f=this.symbolsByDataPath)[b]&&(f[b]=[]),this.symbolsByDataPath[b].push(a);for(i=0,k=e.length;k>i;i++)d=e[i],null==(g=this.symbolsBySubViewName)[d]&&(g[d]=[]),this.symbolsBySubViewName[d].push(a);return a},f.prototype.addSymbol=function(a){return this.symbols[a.id]={id:a.id,tagName:"function"==typeof a.getTagName?a.getTagName():void 0}},f.prototype.appendChild=function(a){return this.addAdhocSymbol(a)},f.prototype.prependChild=function(a){return this.addAdhocSymbol(a)},f.prototype.registerDataPaths=function(a){var b,c,d,e;for(e=[],c=0,d=a.length;d>c;c++)b=a[c],e.push(this.dataPaths[b]=!0);return e},f.prototype.registerSubViewNames=function(a){var b,c,d,e;for(e=[],c=0,d=a.length;d>c;c++)b=a[c],e.push(this.subViewNames[b]=!0);return e},f.prototype.getDataPaths=function(){return Object.keys(this.dataPaths)},f.prototype.getSubViewNames=function(){return Object.keys(this.subViewNames)},a=function(a){return a.replace(/(this\["|\"])/g,"")},e={subview:"symbolsBySubViewName",path:"symbolsByDataPath"},f.prototype.refreshChildren=function(a,b,d){var f,g,h,i,j,k,l,m,n,o,p,q;for(null==d&&(d=function(){}),l={},m=0,o=b.length;o>m;m++)if(h=b[m],k=this[e[a]][h],null!=k)for(n=0,p=k.length;p>n;n++)j=k[n],l[j.id]=j;q=[];for(g in l)c.call(l,g)&&(j=l[g],f=this.view.getElement().querySelector("#"+g),null!=f&&(i=null!=j?j.render():void 0,i?q.push(d.call(f,i)):q.push(void 0)));return q},f.prototype.embedSubViews=function(a){return null==a&&(a=this.getSubViewNames()),this.refreshChildren("subview",a)},f.prototype.update=function(a){return null==a&&(a=this.getDataPaths()),this.refreshChildren("path",a,function(a){return this.innerHTML=a})},f}(),this.Pistachio=a}.call(this);