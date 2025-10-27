interface ProductComponentsProps{
    type: string,
    id: string | undefined
}

const desc = ("Camisa Tech Feminina  Leve, respirável e de secagem rápida, com modelagem ajustada que valoriza o corpo e garante conforto em qualquer ocasião." +
"Características: Tecido tecnológico de alta performance; Secagem rápida e respirabilidade; Leve, elástico e confortável;" +
"Modelagem feminina levemente ajustada;Ideal para treinos, viagens e uso diário")


export default function ProductComponents(data : ProductComponentsProps) {

    const title: string = "";
    const first_part: string = "";
    const second_title: string = "";
    const second_title: string = "";


    if (desc.indexOf('-') !== -1)
        title = desc.slice(0, desc.indexOf('-'));

    if (desc.indexOf('.') !== -1)
        first_part = (title != undefined) ? desc.indexOf('-') : 0 , desc.indexOf('.'))

    const first_part : string = ( ? desc.slice ( 
    const secound_part : string = (desc.indexOf('.') !== 0) ? "a" : "b"

    if (data.type == 'DSC') {
        return (
            <div className="flex w-75 text-sm font-noto">
                <p > <span className="font-bold">{title} </span>
                    {first_part}
                </p>
            </div>
        )
    }

    return (
        <div>
        </div>
    )
}